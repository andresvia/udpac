# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|

  config.vm.provision "shell", inline: "rm -rf /tmp/packages"
  config.vm.provision "file", source: "packages", destination: "/tmp/packages"

  config.vm.define(:debian) do |debian|
    debian.vm.box = "debian/jessie64"
    debian.vm.provision "shell", inline: "sudo dpkg -i /tmp/packages/*.deb"
    debian.vm.provision "shell", inline: "rm -rf /tmp/packages/"
    debian.vm.network :forwarded_port, guest: 80, host: 9645
  end

  config.vm.define(:centos) do |centos|
    centos.vm.box = "centos/7"
    centos.vm.provision "shell", inline: "sudo rpm -Uhv /tmp/packages/*.rpm"
    centos.vm.provision "shell", inline: "rm -rf /tmp/packages/"
    centos.vm.network :forwarded_port, guest: 80, host: 9646
  end

end
