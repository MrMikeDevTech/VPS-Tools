# VPS Tools

Mi CLI `VPS-Tools` es una **herramienta escrita en Go** diseñada para simplificar la **gestión y administración de mi servidor VPS privado**.  
Está pensada para uso personal y no está diseñada para ser utilizada en entornos públicos o compartidos.

[![Build & Deploy VPS CLI](https://github.com/MrMikeDevTech/VPS-Tools/actions/workflows/pipeline.yml/badge.svg?branch=master)](https://github.com/MrMikeDevTech/VPS-Tools/actions/workflows/pipeline.yml)
[![Beta](https://img.shields.io/badge/status-beta-yellow)](https://github.com/MrMikeDevTech/VPS-Tools)
[![Made with Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?logo=go)](https://go.dev/)
---

![vps-welcome](/public/vps-tools-welcome.png)

## ✨ Características

- 🖥️ Información detallada del sistema con arte ASCII
- ⚡ Binario ligero y sin dependencias externas
- 🐧 Diseñado específicamente para Linux y servidores VPS
- 🔒 Herramientas para configuración de seguridad y firewall
- 🚀 Automatización de despliegues y administración de servicios
- 📂 Soporte para templates de proyectos personalizados
- 🛠️ Compatible con entornos CI/CD mediante modo no-interactivo

---

## 📦 Instalación

### Desde binario
```bash
chmod +x vps
sudo mv vps /usr/local/bin/vps
```

### Desde código fuente
```bash
git clone https://github.com/MrMikeDevTech/vps-tools.git
cd vps-tools
GOOS=linux GOARCH=amd64 go build -o vps
sudo mv vps /usr/local/bin/
```

---

## 🖥️ Uso

```bash
vps --help
```

![vps-help](/public/vps-tools-help.png)

---

## 📊 Información del sistema

```bash
vps fetch
```

Muestra información del sistema con arte ASCII según el sistema operativo, similar a **neofetch**.

![vps-fetch](/public/vps-tools-fetch.png)

---

## 🧠 Tecnologías usadas

- Go
- Cobra
- gopsutil

---

## 🛠️ En desarrollo

- [ ] Configuración de firewall
- [ ] Instalación de servicios
- [ ] Despliege de apps
- [ ] Deploy para Go / Bun
- [ ] Templates de proyectos
- [ ] Logs y monitoreo
- [ ] Autocompletado (bash/zsh)
- [ ] Modo no-interactivo (CI/CD)

---

## 📜 Licencia

MIT © MrMikeDev
