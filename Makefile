go_wasm::
	cd tinygo && tinygo build --no-debug -scheduler=none -target=wasi -o filter.wasm ./filter.go && \
	wasm-opt -Oz -o go.wasm filter.wasm

rust_wasm::
	cd rust && cargo build --target wasm32-unknown-unknown --release && \
	wasm-opt -Oz -o rust.wasm target/wasm32-unknown-unknown/release/filter_rust.wasm

rust_fb::
	cd rust && fluent-bit -c fb_config

go_fb::
	cd tinygo && fluent-bit -c fb_config