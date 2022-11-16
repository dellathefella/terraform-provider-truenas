---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "truenas_vm Resource - terraform-provider-truenas"
subcategory: ""
description: |-
  
---

# truenas_vm (Resource)



## Example Usage

```terraform
resource "truenas_vm" "vm" {
  name = "TestVM"
  description = "Test VM"
  vcpus = 2
  bootloader = "UEFI"
  autostart = true
  time = "UTC"
  shutdown_timeout = "10"
  cores = 4
  threads = 2
  memory = 1024*1024*512 // 512MB

  device {
    type = "NIC"
    attributes = {
      type = "VIRTIO"
      mac = "00:a0:98:39:5b:78"
      nic_attach = "br4"
    }
  }

  device {
    type = "DISK"
    attributes = {
      path = "/dev/zvol/Tank/dev-3qsqd"
      type = "AHCI"
      physical_sectorsize = null
      logical_sectorsize = null
    }
  }

  device {
    type = "DISPLAY"
    attributes = {
      wait = false
      port = 9736
      resolution = "1024x768"
      bind = "0.0.0.0"
      password = ""
      web = true
      type = "VNC"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) VM name

### Optional

- `autostart` (Boolean) Set to start this VM when the system boots
- `bootloader` (String) VM bootloader
- `cores` (Number) Specify the number of cores per virtual CPU socket. The product of vCPUs, cores, and threads must not exceed 16.
- `description` (String) VM description
- `device` (Block Set) (see [below for nested schema](#nestedblock--device))
- `memory` (Number) Allocate RAM for the VM. Minimum value is 256 * 1024 * 1024 B. Units are bytes. Allocating too much memory can slow the system or prevent VMs from running
- `shutdown_timeout` (Number) The time in seconds the system waits for the VM to cleanly shut down. During system shutdown, the system initiates poweroff for the VM after the shutdown timeout has expired.
- `threads` (Number) Specify the number of threads per core. The product of vCPUs, cores, and threads must not exceed 16.
- `time` (String) VM system time. Default is `Local`
- `vcpus` (Number) Number of virtual CPUs to allocate to the virtual machine. The maximum is 16, or fewer if the host CPU limits the maximum. The VM operating system might also have operational or licensing restrictions on the number of CPUs.

### Read-Only

- `id` (String) The ID of this resource.
- `status` (Set of Object) (see [below for nested schema](#nestedatt--status))
- `vm_id` (String) VM ID

<a id="nestedblock--device"></a>
### Nested Schema for `device`

Required:

- `attributes` (Map of String) Device attributes specific to device type, check VM resource examples for example device configurations
- `type` (String) Device type

Read-Only:

- `id` (String) Device ID
- `order` (Number) Device order
- `vm` (Number) Device VM ID


<a id="nestedatt--status"></a>
### Nested Schema for `status`

Read-Only:

- `domain_state` (String)
- `pid` (Number)
- `state` (String)

