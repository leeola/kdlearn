---
author: William
username: willwac
date: 2014-12-20
description: 'Koding with Sublime Text'
categories: [Sublime Text]
---

# Koding with Sublime Text

This guide will show you how to connect your Koding VM with Sublime Text via SFTP.

## What you will need

- Set up FTP on your [Koding](https://koding.com) account: http://learn.koding.com/guides/setting-up-ftp-on-koding
- Your [Koding](https://koding.com) username
- Your [Koding](https://koding.com) VM Hostname
- [Sublime Text](https://www.sublimetext.com/) 2 or 3

## Getting Down to Business

1. Download [Package Control](https://sublime.wbond.net/installation#st2) and use the Package Control: Install Package command from the command palette. The command palette is opened via ctrl+shift+p on Windows and Linux, and cmd+shift+p on OS X.

  ![alt tag](st1.png)

2. Once you run the Package Control: Install Package command, a list of packages will appear. Type in SFTP and hit enter. Using Package Control ensures Sublime SFTP will stay up to date automatically.

  ![alt tag](st2.png)

3. Open the folder that you want to sync with your VM using Sublime Text

  ![alt tag](st4.png)

4. Right click the folder in the sidebar and select `SFTP/FTP` -> `Map to Remote`

  ![alt tag](st5.png)

5. You should now see `sftp-config.json` opened.

  ![alt tag](st6.png)

6. Configure your `sftp-config.json` file and save it.

  ```
  // sftp, ftp or ftps
    "type": "ftp",
  "host": "<hostname>.<username>.koding.io",
  "user": "<username>",
  "password": "<koding_password>",
  //"port": "22",
    
  "remote_path": "/home/<username>/",
  ```

  ![alt tag](st7.png)

   Modify the other configurations if you wish. I personally find the option to `upload_on_save` very handy.
   
  ![alt tag](st8.png)

7. Right click the folder, select `SFTP/FTP`. You will now see many more options than before.

  ![alt tag](st9.png)

8. (**IMPORTANT**) Chose whether you want to `Sync Local -> Remote` or `Sync Remote -> Local`. Be very careful about which you choose.

   `Sync Local -> Remote` overwrites what you have in your remote directory that you specified in you `sftp-config.json` file with what you have in your local directory.
   
   `Sync Remote -> Local` overwrites what you have in your local directory with what you have in your remote directory that you specified in you `sftp-config.json` file.
   
  ![alt tag](st10.png)

9. (**IMPORTANT**) Check what is about to be performed in the box below before confirming your actions!

  ![alt tag](st11.png)

10. And there you have it, a folder synced with your Koding VM.

  ![alt tag](st12.png)
