all:
  children:
    app:
      hosts:
        ${app_ip}:
      vars:
        ansible_user: root
        ansible_ssh_private_key_file: ~/.ssh/id_rsa
