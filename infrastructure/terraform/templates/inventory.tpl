all:
  children:
    loadbalancer:
      hosts:
        ${lb_ip}:
          ansible_host: ${lb_ip}
      vars:
        ansible_user: root
        ansible_ssh_private_key_file: ~/.ssh/id_rsa
    app:
      hosts:
        ${app_ip}:
          ansible_host: ${app_ip}
      vars:
        ansible_user: root
        ansible_ssh_private_key_file: ~/.ssh/id_rsa
