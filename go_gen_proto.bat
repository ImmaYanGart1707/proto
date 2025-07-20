@echo off
setlocal enabledelayedexpansion

echo Cleaning up existing .pb.go files...
echo.

if exist "../gen" (
    cd ../gen
    echo Working in proto directory...
) else (
    echo Proto directory not found, working in current directory...
)

set deletecount=0

for /r %%f in (*.pb.go) do (
    set /a deletecount+=1
    set "filepath=%%f"
    set "relativepath=!filepath:%CD%\=!"
    echo [!deletecount!] Deleting: !relativepath!
    del "%%f"
    if !errorlevel! equ 0 (
        echo     ✓ Deleted successfully
    ) else (
        echo     ✗ Failed to delete with error code !errorlevel!
    )
)

if !deletecount! equ 0 (
    echo No existing .pb.go files found.
) else (
    echo Deleted !deletecount! .pb.go files.
)

echo Starting protoc generation for all .proto files...
echo.

if exist "../proto" (
    cd ../proto
    echo Working in proto directory...
) else (
    echo Proto directory not found, working in current directory...
)

echo.
echo ----------------------------------------
echo Starting .proto files generation...
echo.

set count=0

for /r %%f in (*.proto) do (
    set /a count+=1

    set "filepath=%%f"
    set "relativepath=!filepath:%CD%\=!"

    echo [!count!] Processing: !relativepath!
    protoc -I proto "!relativepath!" --go_out=gen\go --go_opt=paths=source_relative --go-grpc_out=gen\go --go-grpc_opt=paths=source_relative

    if !errorlevel! equ 0 (
        echo     ✓ Success
    ) else (
        echo     ✗ Failed with error code !errorlevel!
    )
    echo.
)

echo.
echo Completed! Deleted !deletecount! old files and processed !count! .proto files.