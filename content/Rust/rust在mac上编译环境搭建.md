rust 在mac机上编译的环境搭建
------------------------------
1. 安装rust
2. 安装lipo: `cargo install cargo-lipo`
3. 安装cbindgen: `cargo install cbindgen`

编译:
````
cargo lipo --release
cbindgen ./src/lib.rs -l c > lib.h
````