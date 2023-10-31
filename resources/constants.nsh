; Define constants with !ifndef

!ifndef VERSION
    !define VERSION "1.0.0.0"
!endif

!ifndef BIN_NAME
    !define BIN_NAME "tarkov-server-checker.exe"
!endif

!ifndef BUILD_PATH
    !define BUILD_PATH ".\bin\tarkov-server-checker.exe"
!endif

!ifndef INSTALLER_BIN_NAME
    !define INSTALLER_BIN_NAME "tarkov-server-checker-setup.exe"
!endif

!ifndef INSTALLER_BUILD_PATH
    !define INSTALLER_BUILD_PATH ".\bin\${INSTALLER_BIN_NAME}"
!endif



