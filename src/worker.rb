require "ffi"

module Worker
  extend FFI::Library
  ffi_lib "./lib/worker.so"

  attach_function :login, [:string], :string
end
