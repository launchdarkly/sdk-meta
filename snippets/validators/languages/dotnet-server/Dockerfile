FROM mcr.microsoft.com/dotnet/sdk:8.0

# python3 is used by the harness's USING_LIFT_MARKER pre-step
# (see run.sh) to hoist `using …;` lines out of doc-fragment method
# bodies before the C# compiler runs.
RUN apt-get update && apt-get install -y --no-install-recommends \
        python3 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /work

COPY shared /harness-shared
COPY languages/dotnet-server/harness /harness

ENTRYPOINT ["/harness/run.sh"]
