# -*- mode: ruby -*-
# vi: set ft=ruby :

VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  #Set box
  config.vm.box = "ubuntu/trusty64"

  #Do not check fo updates
  config.vm.box_check_update = false

  #Add private network
  config.vm.network "private_network", ip: "192.168.192.168"

  #Sync folder using NFS if not windows
  config.vm.synced_folder ".", "/vagrant",
    :nfs => !Vagrant::Util::Platform.windows?

  #Virtal Box settings
  config.vm.provider "virtualbox" do |vb|
    # Don't boot with headless mode
    #vb.gui = true

    # Set VM settings
    vb.customize ["modifyvm", :id, "--memory", "2048"]
    vb.customize ["modifyvm", :id, "--cpus", 2]
  end

  #Install dependencies
  config.vm.provision "shell",
      inline: "sudo sh /vagrant/vm-init/install.sh"
end
