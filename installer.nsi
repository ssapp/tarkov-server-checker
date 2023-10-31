!include "LogicLib.nsh"
!include "x64.nsh"
!include "MUI.nsh"
!include "nsDialogs.nsh"

!define APP_NAME "Tarkov Server Checker"
!define COMP_NAME "https://github.com/ssapp/tarkov-server-checker"
!define WEB_SITE "https://github.com/ssapp/tarkov-server-checker"
!define COPYRIGHT "Vladislav Naydenov (c) 2023"
!define DESCRIPTION "Escape from Tarkov Server Location Checker"
!define INSTALL_TYPE "SetShellVarContext current"
!define REG_ROOT "HKCU"
!define REG_APP_PATH "Software\Microsoft\Windows\CurrentVersion\App Paths\${BIN_NAME}"
!define UNINSTALL_PATH "Software\Microsoft\Windows\CurrentVersion\Uninstall\${APP_NAME}"
!define MUI_ICON "${ICON_PATH}"
!define MUI_UNICON "${ICON_PATH}"
var SM_Folder

;--------------------------------
; Interface Configuration

!define MUI_ABORTWARNING
!define MUI_UNABORTWARNING

;--------------------------------
; The stuff to install
VIProductVersion "${VERSION}"
VIAddVersionKey "ProductName" "${APP_NAME}"

;------------------------------------------------
; General
Name "${APP_NAME}"
Caption "${APP_NAME}"
OutFile "${INSTALLER_BUILD_PATH}"
BrandingText "${APP_NAME}"
InstallDirRegKey "${REG_ROOT}" "${REG_APP_PATH}" ""
InstallDir "$PROGRAMFILES64\${APP_NAME}"
ShowInstDetails hide
XPStyle on

Var Checkbox


;--------------------------------
; Interface Settings
!insertmacro MUI_PAGE_DIRECTORY

!insertmacro MUI_PAGE_LICENSE "LICENSE.txt"

!ifndef REG_START_MENU
    !insertmacro MUI_PAGE_STARTMENU Application $SM_Folder
!endif

!insertmacro MUI_PAGE_INSTFILES
!define MUI_FINISHPAGE_RUN "$INSTDIR\${BIN_NAME}"
!insertmacro MUI_PAGE_FINISH
!define MUI_FINISHPAGE_RUN_FUNCTION "MyFunction"
!define MUI_FINISHPAGE_NOAUTOCLOSE false

!insertmacro MUI_UNPAGE_WELCOME
!insertmacro MUI_UNPAGE_CONFIRM
!insertmacro MUI_UNPAGE_INSTFILES
!insertmacro MUI_UNPAGE_FINISH 

;------------------------------------------------
; Languages
!insertmacro MUI_LANGUAGE "English"

;--------------------------------
; Installer Sections
Section -MainProgram
    ${INSTALL_TYPE}
    SetOverwrite ifnewer
    SetOutPath "$INSTDIR"
    File "${BUILD_PATH}"
SectionEnd

;--------------------------------
; Uninstaller Sections
Section -Icons_Reg
SetOutPath "$INSTDIR"
WriteUninstaller "$INSTDIR\uninstall.exe"
CreateDirectory "$SMPROGRAMS\$SM_Folder"
CreateShortCut "$SMPROGRAMS\$SM_Folder\${APP_NAME}.lnk" "$INSTDIR\${BIN_NAME}"
CreateShortCut "$DESKTOP\${APP_NAME}.lnk" "$INSTDIR\${BIN_NAME}"
CreateShortCut "$SMPROGRAMS\$SM_Folder\Uninstall ${APP_NAME}.lnk" "$INSTDIR\uninstall.exe"
WriteIniStr "$INSTDIR\${APP_NAME} website.url" "InternetShortcut" "URL" "${WEB_SITE}"
CreateShortCut "$SMPROGRAMS\$SM_Folder\${APP_NAME} Website.lnk" "$INSTDIR\${APP_NAME} website.url"

WriteRegStr ${REG_ROOT} "${REG_APP_PATH}" "" "$INSTDIR\${BIN_NAME}"
WriteRegStr ${REG_ROOT} "${UNINSTALL_PATH}" "DisplayName" "${APP_NAME}"
WriteRegStr ${REG_ROOT} "${UNINSTALL_PATH}" "UninstallString" "$INSTDIR\uninstall.exe"
WriteRegStr ${REG_ROOT} "${UNINSTALL_PATH}" "DisplayIcon" "$INSTDIR\${BIN_NAME}"
WriteRegStr ${REG_ROOT} "${UNINSTALL_PATH}" "DisplayVersion" "${VERSION}"
WriteRegStr ${REG_ROOT} "${UNINSTALL_PATH}" "Publisher" "${COMP_NAME}"
WriteRegStr ${REG_ROOT} "${UNINSTALL_PATH}" "URLInfoAbout" "${WEB_SITE}"
SectionEnd

;--------------------------------
; Descriptions
Section Uninstall
    ${INSTALL_TYPE}
    Delete "$INSTDIR\tarkov-server-checker.exe"
    Delete "$INSTDIR\uninstall.exe"
    Delete "$INSTDIR\${APP_NAME} website.url"
    RmDir "$INSTDIR"
    Delete "$SMPROGRAMS\$SM_Folder\${APP_NAME}.lnk"
    Delete "$SMPROGRAMS\$SM_Folder\Uninstall ${APP_NAME}.lnk"
    Delete "$SMPROGRAMS\$SM_Folder\${APP_NAME} Website.lnk"
    Delete "$DESKTOP\${APP_NAME}.lnk"
    RmDir "$SMPROGRAMS\$SM_Folder"
    DeleteRegKey ${REG_ROOT} "${REG_APP_PATH}"
    DeleteRegKey ${REG_ROOT} "${UNINSTALL_PATH}"
SectionEnd

;--------------------------------
; Macro to check if already installed
!macro checkAlreadyInstalled
    ReadRegStr $0 ${REG_ROOT} "${REG_APP_PATH}" ""
    StrCmp $0 "" installed
    ${if} "${INSTALL_TYPE}" == "SetShellVarContext current"
        StrCmp $0 "$INSTDIR" installed
    ${else}
        StrCmp $0 "$PROGRAMFILES64\Tarkov Server Checker" installed
    ${endif}
    MessageBox MB_OKCANCEL \
        "${APP_NAME} is already installed!$\n$\nDo you want to reinstall it?" \
        IDCANCEL installed
        ExecWait '"$INSTDIR\uninstall.exe"'
        Quit
    installed:
!macroend

!macro asd 
    nsExec::ExecToStack "tasklist /FI \"IMAGENAME eq ${BIN_NAME}\""
    Pop $0
    StrCmp $0 "" keepRunning
    MessageBox MB_OKCANCEL \
        "${APP_NAME} is running!$\n$\nDo you want to terminate this process and proceed with the upgrade/reinstall?" \
        IDCANCEL keepRunning
        nsExec::ExecToStack "taskkill /IM ${BIN_NAME} /F"
        Quit
    keepRunning:
!macroend

;--------------------------------
; Installer functions
Function MyFunction
    ExecShell "" "$INSTDIR\${BIN_NAME}"
FunctionEnd

Function .onInit
    !insertmacro checkAlreadyInstalled
FunctionEnd

Function .onInstFailed
    !insertmacro asd
FunctionEnd

Function un.onInit
FunctionEnd

;--------------------------------
; Macro to check if running and kill
!macro terminateProcess
    nsExec::ExecToStack "taskkill /IM ${BIN_NAME} /F"
    Pop $0
!macroend

;--------------------------------
; Uninstaller functions
Function un.onUninstSuccess
    !insertmacro terminateProcess
FunctionEnd

;--------------------------------
; Installer attributes
!insertmacro MUI_FUNCTION_DESCRIPTION_BEGIN
!insertmacro MUI_DESCRIPTION_TEXT ${SM_Folder} "Start Menu Folder"
!insertmacro MUI_FUNCTION_DESCRIPTION_END
