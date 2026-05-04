# devmapper-demo

A Rust program that demonstrates Linux device mapper thin provisioning and loop devices. It builds a working thin pool from scratch, 
creates volumes on top of it, takes snapshots, and shows copy-on-write in action. No external Rust crates beyond `anyhow` and `nix` are needed for the core logic. 
The heavy lifting is done entirely through Linux kernel interfaces.

## What it does

The program sets up two loop devices backed by sparse files in `/tmp`, one for pool metadata and one for actual data storage. 
It initializes the metadata device using `thin_restore` from the thin-provisioning-tools suite, writing a minimal XML superblock 
that describes an empty pool structure before the kernel ever touches it. On top of those two loop devices, it creates a device mapper thin pool target using `dmsetup`.

From there it provisions a 500MB thin volume from the pool, formats it with ext4, mounts it, and writes two files to it. 
It then takes a snapshot of that base volume and mounts the snapshot separately. When the snapshot's copy of `shared-file.txt` is overwritten, 
the kernel silently copies the affected data blocks before writing, leaving the base volume's copy completely intact. The program verifies this explicitly before exiting.

All three types implement `Drop`, so the pool, volumes, and loop devices detach and remove themselves automatically even if the program exits mid-run due to an error.

The `demonstrate_loop_device` function is also present in the source but commented out in `main`. You can uncomment it to run a standalone demo that just creates a loop device, 
inspects it with `lsblk` and `losetup`, then cleans up.

## Prerequisites

You need to be running Linux with the following tools available on your PATH

- `fallocate`
- `losetup`
- `blockdev`
- `dmsetup`
- `thin_restore` from the `thin-provisioning-tools` package
- `mkfs.ext4`

Install thin-provisioning-tools on Ubuntu or Debian

```
sudo apt install thin-provisioning-tools e2fsprogs
```

On Fedora or RHEL-based systems

```
sudo dnf install device-mapper-persistent-data e2fsprogs
```

## Running it

The program needs root to manage loop devices, mount filesystems, and interact with device mapper targets. Run it after building with cargo

```
cargo build --release
sudo ./target/release/thin-provisioning-demo
```

Expected output walks through each stage, backing file allocation, loop device attachment, metadata initialization, pool creation, volume formatting, 
snapshot creation, and the final verification that the base volume is unchanged. At the end, `dmsetup status` output for the pool is printed so you can see actual block allocation numbers.

## What to watch for

When the snapshot modifies `shared-file.txt`, the kernel triggers a copy-on-write at the data block level, not the file level. The pool's data chunk size is set to 1MB (2048 sectors of 512 bytes each), 
so the smallest unit of copy-on-write is one of those chunks. After the snapshot write, if you run `dmsetup status demo-pool` manually before the program exits, 
you will see the used data block count increase by at least one chunk relative to what it was before the write.

## Project layout

```
src/
  main.rs   -- everything lives here
Cargo.toml
```

Three main types drive the program. `LoopDevice` wraps a backing file and its `/dev/loopN` attachment and handles both allocation and detachment. 
`ThinPool` owns the two loop devices and the dm thin-pool target sitting on top of them, and exposes methods for provisioning volumes. 
`ThinVolume` represents individual thin volumes and snapshots created from the pool and handles both device removal and the `dmsetup message delete` call needed to release the device ID back to the pool. 
All three implement `Drop` so cleanup is automatic even if something fails mid-run.

## Known limitations

The backing files are created under `/tmp` with hardcoded names. Running two instances simultaneously will conflict. There is also no resume path; 
if the process is killed before cleanup the loop devices and dm targets will remain and need to be removed manually with `losetup -d` and `dmsetup remove`.
