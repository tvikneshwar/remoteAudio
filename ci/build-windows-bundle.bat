mkdir %GOPATH%\src\github.com\dh1tw\remoteAudio\release\
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libogg-0.dll /c/gopath/src/github.com/dh1tw/remoteAudio/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libopus-0.dll /c/gopath/src/github.com/dh1tw/remoteAudio/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libopusfile-0.dll /c/gopath/src/github.com/dh1tw/remoteAudio/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libportaudio-2.dll /c/gopath/src/github.com/dh1tw/remoteAudio/release/"
%MSYS_PATH%\usr\bin\bash -lc "cp /mingw%MSYS2_BITS%/**/libsamplerate-0.dll /c/gopath/src/github.com/dh1tw/remoteAudio/release/"
REM %MSYS_PATH%\usr\bin\bash -lc "pacman --noconfirm --needed -Sy sed" > nul
REM %MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio && ci/release"
%MSYS_PATH%\usr\bin\bash -lc "cp /c/gopath/src/github.com/dh1tw/remoteAudio/remoteAudio.exe /c/gopath/src/github.com/dh1tw/remoteAudio/release"
%MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio/release && 7z a -tzip remoteAudio.zip *"
REM %MSYS_PATH%\usr\bin\bash -lc "cd /c/gopath/src/github.com/dh1tw/remoteAudio && rm -rf ./release"
xcopy %GOPATH%\src\github.com\dh1tw\remoteAudio\release\remoteAudio.zip %APPVEYOR_BUILD_FOLDER%\ /e /i > nul