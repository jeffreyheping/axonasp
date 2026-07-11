#define MyAppName "AxonASP Server"
; Allow CI to override version and arch suffix via iscc /D flag.
#ifndef MyAppVersion
  #define MyAppVersion "2.3.0"
#endif
#define MyAppPublisher "G3pix Ltda"
#define MyAppURL "https://g3pix.com.br/axonasp"
#define MyAppExeName "axonasp-http.exe"
#ifndef MyArchSuffix
  #define MyArchSuffix "amd64"
#endif
#ifndef MyVersionInfoVersion
  #define MyVersionInfoVersion MyAppVersion
#endif

[Setup]
AppId={{0E1F2C1D-3A4B-5C6D-7E8F-901234567890}}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
DefaultDirName={autopf}\AxonASP
DefaultGroupName=AxonASP
AllowNoIcons=yes
LicenseFile=LICENSE.txt
OutputDir=Output
OutputBaseFilename=axonasp-installer-{#MyAppVersion}-{#MyArchSuffix}
Compression=lzma2/ultra
SolidCompression=yes
WizardStyle=modern
ArchitecturesInstallIn64BitMode=x64 arm64
UninstallDisplayIcon={app}\axonasp-http.exe
VersionInfoVersion={#MyVersionInfoVersion}
VersionInfoCompany={#MyAppPublisher}
VersionInfoDescription=AxonASP Server Setup
VersionInfoProductName=AxonASP Server
SetupIconFile=resources\icon_server.ico

[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Tasks]
Name: "desktopicon"; Description: "Create a &desktop shortcut"; GroupDescription: "Additional icons:"

[Files]
Source: "axonasp-http.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "axonasp-fastcgi.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "axonasp-cli.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "axonasp-mcp.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "axonasp-admin.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "axonasp-service.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "axonasp-testsuite.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "www\*"; DestDir: "{app}\www"; Flags: ignoreversion recursesubdirs createallsubdirs
Source: "config\*"; DestDir: "{app}\config"; Flags: ignoreversion recursesubdirs createallsubdirs
Source: "mcp\*"; DestDir: "{app}\mcp"; Flags: ignoreversion recursesubdirs createallsubdirs
Source: "LICENSE.txt"; DestDir: "{app}"; Flags: ignoreversion
Source: "LEGAL-DISCLAIMER.md"; DestDir: "{app}"; Flags: ignoreversion
Source: "global.asa"; DestDir: "{app}"; Flags: ignoreversion
Source: "iis-http.cmd"; DestDir: "{app}"; Flags: ignoreversion
Source: "install-service.ps1"; DestDir: "{app}"; Flags: ignoreversion
Source: "uninstall-service.ps1"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\AxonASP HTTP Server"; Filename: "{app}\{#MyAppExeName}"
Name: "{group}\AxonASP CLI"; Filename: "{app}\axonasp-cli.exe"
Name: "{group}\AxonASP Admin"; Filename: "{app}\axonasp-admin.exe"
Name: "{group}\Uninstall AxonASP"; Filename: "{uninstallexe}"
Name: "{autodesktop}\AxonASP Server"; Filename: "{app}\{#MyAppExeName}"; Tasks: desktopicon

[Registry]
Root: HKLM; Subkey: "SYSTEM\CurrentControlSet\Control\Session Manager\Environment"; ValueType: expandsz; ValueName: "Path"; ValueData: "{olddata};{app}"; Check: NeedsAddPath('{app}'); Flags: uninsdeletevalue

[Code]
function NeedsAddPath(Param: string): boolean;
var
  OrigPath: string;
begin
  if not RegQueryStringValue(HKLM, 'SYSTEM\CurrentControlSet\Control\Session Manager\Environment', 'Path', OrigPath) then begin
    Result := True;
    exit;
  end;
  Result := Pos(';' + Uppercase(Param) + ';', ';' + Uppercase(OrigPath) + ';') = 0;
end;
