set times=12
for /l %%i in (1, 1, %times%) do (
    cd mysql_%%i%
	start cmd /k stop.bat
	cd ..
)
taskkill /f /im cmd.exe