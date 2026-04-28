FROM mcr.microsoft.com/dotnet/sdk:8.0

WORKDIR /work

COPY shared /harness-shared
COPY languages/dotnet-server/harness /harness

ENTRYPOINT ["/harness/run.sh"]
