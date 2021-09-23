This program reads a file with dates and notifies about events that are happening that day (it uses `notify-send` to
send the notifications). I put it in my `.xinitrc`, so it runs on every login. Location of the file with dates can
be `~/.remind_me` or `~/.config/reminder/remind_me`. The program doesn't parse the file, so don't expect error messages
if the file is in a bad format. Example of how the file can look like:
```
21-09-23 : This will be shown on September the 23rd 2021
xx-xx-20 : This will be shown every month on the 20th
xx-xx-x{0,2,4,6,8} : This will be shown every even day
```