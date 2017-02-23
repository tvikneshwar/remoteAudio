mkdir %GOPATH%\src\github.com\dh1tw\
xcopy %APPVEYOR_BUILD_FOLDER%\* %GOPATH%\src\github.com\dh1tw\remoteAudio /e /i /s /EXCLUDE:%MSYS_PATH% > nul
xcopy %APPVEYOR_BUILD_FOLDER%\.git %GOPATH%\src\github.com\dh1tw\remoteAudio\ /e /i /s
dir %GOPATH%\src\github.com\dh1tw\remoteAudio
if "%METHOD%"=="ci" SET MSYS_PATH=c:\msys64
if "%METHOD%"=="cross" SET MSYS_PATH=%APPVEYOR_BUILD_FOLDER%\msys%MSYS2_BITS%
SET PATH=%MSYS_PATH%\usr\bin;%PATH%
SET PATH=%MSYS_PATH%\mingw%MSYS2_BITS%\bin;%PATH%
if "%METHOD%"=="cross" appveyor DownloadFile http://kent.dl.sourceforge.net/project/msys2/Base/%MSYS2_ARCH%/msys2-base-%MSYS2_ARCH%-%MSYS2_BASEVER%.tar.xz
if "%METHOD%"=="cross" 7z x msys2-base-%MSYS2_ARCH%-%MSYS2_BASEVER%.tar.xz > nul
if "%METHOD%"=="cross" 7z x msys2-base-%MSYS2_ARCH%-%MSYS2_BASEVER%.tar > nul
%MSYS_PATH%\usr\bin\bash -lc "echo update-core starting..." 2> nul
%MSYS_PATH%\usr\bin\bash -lc "update-core" > nul
%MSYS_PATH%\usr\bin\bash -lc "echo install-deps starting..."
%MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy mingw-w64-%MSYS2_ARCH%-gcc" > nul
%MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy mingw-w64-%MSYS2_ARCH%-pkg-config" > nul
%MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy mingw-w64-%MSYS2_ARCH%-libsamplerate" > nul
%MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy mingw-w64-%MSYS2_ARCH%-portaudio" > nul
%MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy mingw-w64-%MSYS2_ARCH%-opusfile" > nul
%MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy mingw-w64-%MSYS2_ARCH%-opus" > nul
%MSYS_PATH%\usr\bin\bash -lc "yes|pacman --noconfirm -Sc" > nul
if "%METHOD%"=="cross" %MSYS_PATH%\autorebase.bat > nul