# OSFC2019 coreboot payloads
Coreboot payloads shipped for OSFC 2019 Darper Pro.

Precompiled binaries can be found in `binary/`.
The source files are in `src/`.

# LinuxBoot payload

1. Build the GNU/Linux kernel using the defconfig in `src/LinuxBoot/`.

2. Build u-root instructions:

```
u-root -build=bb -initcmd init -defaultsh elvish -files src/mplayer:bin/mplayer -files video.webm:video.webm -o initramfs_u-root.cpio coreboot-app systemboot bootanimation
xz --keep --force --check=crc32 --lzma2=dict=1MiB initramfs_u-root.cpio
```

3. Replace the payload in the vendor firmware:
```
cbfstool firmware.rom remove -n fallback/payload
cbfstool firmware.rom add-payload -f bzImage -n fallback/payload -I initramfs_u-root.cpio.xz
```

# SeaBIOS payload

1. Build SeaBios using the .config in `src/Seabios/`.

2. Replace the payload in the vendor firmware:
```
cbfstool firmware.rom remove -n fallback/payload
cbfstool firmware.rom add -f vgabios.bin -n seavgabios.bin -t raw -r COREBOOT
cbfstool firmware.rom add-payload -f bios.bin.elf -n fallback/payload
```
