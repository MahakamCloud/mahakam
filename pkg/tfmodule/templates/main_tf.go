package templates

// MainFile represents terraform main.tf go template
var MainFile = `module "{{.Name}}" {
    source = "{{.LibvirtModulePath}}"

    instance_name = "${var.hostname}"
    libvirt_host  = "${var.host}"
    source_path   = "${var.image_source_path}"
    mac_address   = "${var.mac_address}"
    ip_address    = "${var.ip_address}"

    user_data = "${data.template_file.user_data.rendered}"
}`
