@echo off
set SOURCE_FOLDER=.
set COMPILER_PATH=.\protoc.exe
set TARGET_PATH=..\server\Source\Msg
set GRPC_PLUGIN_PATH=.\grpc_cpp_plugin.exe
for /f "delims=" %%i in ('dir /b "%SOURCE_FOLDER%\*.rpc"') do (
	echo %COMPILER_PATH% --cpp_out=%TARGET_PATH% %%i
    %COMPILER_PATH% --cpp_out=%TARGET_PATH% %%i
    
	echo %COMPILER_PATH% --grpc_out=%TARGET_PATH% --plugin=protoc-gen-grpc=%GRPC_PLUGIN_PATH% %%i
	%COMPILER_PATH% --grpc_out=%TARGET_PATH% --plugin=protoc-gen-grpc=%GRPC_PLUGIN_PATH% %%i
)

pause