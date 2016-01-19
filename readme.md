lstclone.exe
============

`lstclone.exe` makes the batch-file which removes duplicate files.

```
[C:] lstclone "C:\Users\hayama\Google ドライブ" > t.cmd
```

__Now, lstclone.exe compares files by name and size only__

t.cmd
-----

```
rem "C:\Users\hayama\Google ドライブ\2011\20110202_神田明神\2011020216570000.JPG"
del "C:\Users\hayama\Google ドライブ\Google フォト\2011\20110202_神田明神\2011020216570000.JPG"
del "C:\Users\hayama\Google ドライブ\Pictures\2011\20110202_神田明神\2011020216570000.JPG"

rem "C:\Users\hayama\Google ドライブ\Backup\SAKURA\nyaos.org\d\index.d\8223030343e21313e2330392"
del "C:\Users\hayama\Google ドライブ\Backup\rackbox\nyaos.org\d\index.d\8223030343e21313e2330392"
del "C:\Users\hayama\Google ドライブ\FromOneDrive\gdrive\Backup\SAKURA\nyaos.org\d\index.d\8223030343e21313e2330392"
    :
```




