FROM mcr.microsoft.com/dotnet/sdk:5.0

WORKDIR /tests
COPY *.csproj .
RUN dotnet restore 

COPY . .
RUN dotnet build

ENTRYPOINT dotnet test bin/Debug/net5.0/jtest.dll -l "console;verbosity=detailed"

