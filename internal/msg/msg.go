package msg

type MSG_Template struct {
	Welcome             string
	WatchingChanges     string
	ChangeDetectedIn    string
	ExecCommand         string
	MsgError            string
	ContentOfFileChange string
	ServerRunningAt	 string
	SusscessfulCompilation string
	OutPutCommand string
}

var MSG_Spanish = MSG_Template{
	Welcome:             "Iniciando Builder...",
	WatchingChanges:     "Observando cambios en el directorio:",
	ChangeDetectedIn:    "Cambio detectado en:",
	ExecCommand:         "ejecutando el comando:",
	MsgError:            "Error ejecutando el comando:",
	ContentOfFileChange: "Contenido del archivo cambiado",
	ServerRunningAt:	 "Servidor corriendo en el puerto:",
	SusscessfulCompilation: "Compilación exitosa",
	OutPutCommand: "Salida del comando:",
}

var MSG_English = MSG_Template{
	Welcome:             "Starting Builder...",
	WatchingChanges:     "Watching changes in the directory:",
	ChangeDetectedIn:    "Change detected in:",
	ExecCommand:         "Executing command:",
	MsgError:            "Error executing the command:",
	ContentOfFileChange: "Content of the changed file",
	ServerRunningAt:     "Server running on port:",
	SusscessfulCompilation: "Successful compilation",
	OutPutCommand:       "Command output:",
}

var MSG_Chinese = MSG_Template{
	Welcome:             "启动构建器...",
	WatchingChanges:     "正在监视目录中的更改:",
	ChangeDetectedIn:    "检测到更改:",
	ExecCommand:         "正在执行命令:",
	MsgError:            "执行命令时出错:",
	ContentOfFileChange: "已更改文件的内容",
	ServerRunningAt:     "服务器运行在端口:",
	SusscessfulCompilation: "编译成功",
	OutPutCommand:       "命令输出:",
}

var MSG_French = MSG_Template{
	Welcome:             "Démarrage du Builder...",
	WatchingChanges:     "Surveillance des modifications dans le répertoire:",
	ChangeDetectedIn:    "Changement détecté dans:",
	ExecCommand:         "Exécution de la commande:",
	MsgError:            "Erreur lors de l'exécution de la commande:",
	ContentOfFileChange: "Contenu du fichier modifié",
	ServerRunningAt:     "Serveur en cours d'exécution sur le port:",
	SusscessfulCompilation: "Compilation réussie",
	OutPutCommand:       "Sortie de la commande:",
}

var MSG_German = MSG_Template{
	Welcome:             "Builder wird gestartet...",
	WatchingChanges:     "Änderungen im Verzeichnis werden überwacht:",
	ChangeDetectedIn:    "Änderung erkannt in:",
	ExecCommand:         "Befehl wird ausgeführt:",
	MsgError:            "Fehler beim Ausführen des Befehls:",
	ContentOfFileChange: "Inhalt der geänderten Datei",
	ServerRunningAt:     "Server läuft auf Port:",
	SusscessfulCompilation: "Erfolgreiche Kompilierung",
	OutPutCommand:       "Befehlsausgabe:",
}

var MSG = MSG_English


