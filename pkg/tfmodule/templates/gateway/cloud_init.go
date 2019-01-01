package gateway

var CloudInit = `#cloud-config
password: ${password}
chpasswd: { expire: False }
ssh_pwauth: True
hostname: ${hostname}
fqdn: ${hostname}.${dns_zone_name}
ssh_authorized_keys:
  - ${ssh_public_key}

resolv_conf:
  nameservers: [${dns_address}]
  searchdomains:
    - ${dns_zone_name}

package_upgrade: true

packages:
  - curl
  - apt-transport-https
  - ca-certificates

write_files:
  - path: /opt/cloud-init/setup-network.sh
    permissions: 0644
    content: |
      cat <<EOF >/etc/network/interfaces
      auto lo
      iface lo inet loopback

      auto ens4
      iface ens4 inet static
        address ${public_ip_address}
        netmask ${public_netmask}
        gateway ${public_gateway_ip}
        dns-nameservers ${dns_address}

      auto ens3
      iface ens3 inet static
        address ${ip_address}
        netmask ${netmask}
      EOF

      ifdown ens3 && ifup ens3
      ifdown ens4 && ifup ens4
      systemctl restart networking
      sysctl -w net.ipv4.ip_forward=1

  - path: /opt/cloud-init/setup-iptables.sh
    permissions: 0644
    content: |
      iptables -P INPUT DROP
      iptables -P FORWARD DROP
      iptables -P OUTPUT DROP
      iptables -F
      iptables -F -t nat

      # Allow localhost traffic
      iptables -A INPUT -i lo -j ACCEPT
      iptables -A OUTPUT -o lo -j ACCEPT

      # Allow all outbound traffic
      iptables -A INPUT -m state --state RELATED,ESTABLISHED -j ACCEPT
      iptables -A OUTPUT -j ACCEPT

      # Allow inbound SSH from private network
      iptables -A INPUT -i ens3 -p tcp -m tcp --dport 22 -m state --state NEW,ESTABLISHED -j ACCEPT
      iptables -A OUTPUT -o ens3 -p tcp -m tcp --sport 22 -m state --state ESTABLISHED -j ACCEPT

      # Allow inbound HTTP traffic
      iptables -A INPUT -p tcp -m tcp --dport 80 -m state --state NEW,ESTABLISHED -j ACCEPT
      iptables -A OUTPUT -p tcp -m tcp --sport 80 -m state --state ESTABLISHED -j ACCEPT

      # Allow inbound HTTPS traffic
      iptables -A INPUT -p tcp -m tcp --dport 443 -m state --state NEW,ESTABLISHED -j ACCEPT
      iptables -A OUTPUT -p tcp -m tcp --sport 443 -m state --state ESTABLISHED -j ACCEPT

      # Allow inbound ICMP packet from private network
      iptables -A INPUT -i ens3 -p icmp -m state --state NEW,RELATED,ESTABLISHED -j ACCEPT
      iptables -A OUTPUT -o ens3 -p icmp -m state --state NEW,RELATED,ESTABLISHED -j ACCEPT

      # Allow IP forward from private network to public
      iptables -A FORWARD -i ens4 -o ens3 -m state --state RELATED,ESTABLISHED -j ACCEPT
      iptables -A FORWARD -i ens3 -o ens4 -j ACCEPT

      # Enable SNAT to let VM network access internet
      iptables -t nat -A POSTROUTING -s ${network_cidr} ! -d ${network_cidr} -j MASQUERADE

runcmd:
  - echo "Configuring Gateway VM"
  - bash -ex /opt/cloud-init/setup-network.sh
  - bash -ex /opt/cloud-init/setup-iptables.sh

final_message: "The system is finally up, after $UPTIME seconds"
`
