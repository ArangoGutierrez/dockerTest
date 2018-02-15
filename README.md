# dockerTest

Quick attempt in Go to be able to pull docker containers and run things inside them

This is mostly lizrice/containers-from-scratch, calling some stuff from containers/image and opencontainers/image-tools

Must be run as root, and doesn't clean up /tmp/ after itself.

```
$ echo "Cows for Dave" | sudo ./dockerTester run docker://chuanwen/cowsay /usr/games/cowsay
Fetching container [docker://chuanwen/cowsay /usr/games/cowsay] 
Files will be fetched to /tmp/dockertest401205044 
Getting image source signatures
Copying blob sha256:c954d15f947c57e059f67a156ff2e4c36f4f3e59b37467ff865214a88ebc54d6
 69.57 MB / 69.57 MB [=====================================================] 10s
Copying blob sha256:c3688624ef2b94ab3981564e23e1f48df8f1b988519373ccfb79d7974017cb85
 70.95 KB / 70.95 KB [======================================================] 0s
Copying blob sha256:848fe4263b3b44987f0eacdb2fc0469ae6ff04b2311e759985dfd27ae5d3641d
 629 B / 629 B [============================================================] 0s
Copying blob sha256:23b4459d3b04aa0bc7cb7f7021e4d7bbb5e87aa74a6a5f57475a0e8badbd9a26
 851 B / 851 B [============================================================] 0s
Copying blob sha256:36ab3b56c8f1a3188464886cbe41f42a969e6f9374e040f13803d796ed27b0ec
 164 B / 164 B [============================================================] 0s
Copying blob sha256:10a876940ddfc7e8cc35113245e8700424b47eb902f8bac2204c1b84631fadf4
 2.12 MB / 2.12 MB [========================================================] 0s
Copying config sha256:c7cf192cd3a2e4ea235ffdbd53da2095a68e3d61a93fce17e534688ad88df42f
 3.41 KB / 3.41 KB [========================================================] 0s
Writing manifest to image destination
Storing signatures
Container will be unpacked to /tmp/dockertest401205044 
Running [/usr/games/cowsay] 
Running [/usr/games/cowsay] 
In container at /tmp/dockertest377924956 
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "en_US.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
 _______________
< Cows for Dave >
 ---------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||

```
