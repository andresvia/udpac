# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.provision "shell", inline: "mkdir -p /tmp/udpac"
  config.vm.define(:deb) do |debian|
    debian.vm.box = "debian/jessie64"
    debian.vm.provision "file", source: "*.deb", destination: "/tmp/udpac"
    debian.vm.provision "shell", inline: "sudo dpkg -i /tmp/udpac/*.deb"
    debian.vm.provision "shell", inline: "rm -rf /tmp/udpac"
    debian.vm.network :forwarded_port, guest: 80

  end
  config.vm.define(:rpm) do |centos|
    centos.vm.box = "centos/7"
    centos.vm.provision "file", source: "*.rpm", destination: "/tmp/udpac"
    centos.vm.provision "shell", inline: "sudo rpm -Uhv /tmp/udpac/*.rpm"
    centos.vm.provision "shell", inline: "rm -rf /tmp/udpac"
    centos.vm.network :forwarded_port, guest: 80
  end
end
