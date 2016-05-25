# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.provision "shell", inline: "mkdir -p /tmp/udpac"
  config.vm.provision "file", source: "packages", destination: "/tmp/udpac/"
  config.vm.define(:deb) do |debian|
    debian.vm.box = "debian/jessie64"
    debian.vm.provision "shell", inline: "sudo dpkg -i /tmp/udpac/packages/*.deb"
    debian.vm.provision "shell", inline: "rm -rf /tmp/udpac/"
    debian.vm.network :forwarded_port, guest: 80, host: 9645

  end
  config.vm.define(:rpm) do |centos|
    centos.vm.box = "centos/7"
    centos.vm.provision "shell", inline: "sudo rpm -Uhv /tmp/udpac/packages/*.rpm"
    centos.vm.provision "shell", inline: "rm -rf /tmp/udpac/"
    centos.vm.network :forwarded_port, guest: 80, host: 9646
  end
end
