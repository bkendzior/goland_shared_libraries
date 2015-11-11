require 'ffi'

module GoFuncs
  extend FFI::Library
  ffi_lib './goCallHTTP.so'
  attach_function :GoHTTP, [:string], :string
end

puts GoFuncs.GoHTTP("https://www.google.com/")
