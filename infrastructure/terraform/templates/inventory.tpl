all:
  vars:
    ansible_user: root
    ansible_ssh_private_key_file: ~/.ssh/id_rsa
    domain: enkya.org
    subdomain: playground
    certbot_email: info@enkya.org
  children:
    loadbalancer:
      hosts:
        ${lb_ip}:
          ansible_host: ${lb_ip}
    app:
      hosts:
        ${app_ip}:
          ansible_host: ${app_ip}
