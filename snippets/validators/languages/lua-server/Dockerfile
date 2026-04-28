FROM ubuntu:24.04

# The Lua server SDK is a thin Lua wrapper over the C++ Server SDK's C
# binding, so the validator image needs the C++ toolchain + the SDK's
# transitive deps + lua + luarocks.
RUN apt-get update && apt-get install -y --no-install-recommends \
        build-essential cmake ninja-build pkg-config \
        libboost-all-dev libssl-dev \
        lua5.3 liblua5.3-dev luarocks \
        git ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Build the C++ Server SDK as a shared library and install to a known
# prefix; the lua rockspec's external_dependencies look up the C header
# and shared lib from $LD_DIR/include and $LD_DIR/lib respectively.
ARG CPP_SDKS_REF=launchdarkly-cpp-server-v3.10.1
RUN git clone --depth 1 --branch "${CPP_SDKS_REF}" \
        https://github.com/launchdarkly/cpp-sdks.git /opt/cpp-sdks \
    && mkdir /opt/cpp-sdks/build && cd /opt/cpp-sdks/build \
    && cmake -G Ninja \
             -DBUILD_TESTING=OFF \
             -DCMAKE_BUILD_TYPE=Release \
             -DLD_BUILD_SHARED_LIBS=On \
             -DCMAKE_INSTALL_PREFIX=/opt/cpp-sdks/install .. \
    && cmake --build . --target launchdarkly-cpp-server \
    && cmake --install . \
    && cd / && rm -rf /opt/cpp-sdks/build /opt/cpp-sdks/.git

ENV LD_LIBRARY_PATH=/opt/cpp-sdks/install/lib

# Install the Lua server SDK rock against the built C++ shared lib.
ARG LUA_SDK_VERSION=2.1.3-0
RUN luarocks install --tree=/opt/lua-rocks \
        launchdarkly-server-sdk ${LUA_SDK_VERSION} \
        LD_DIR=/opt/cpp-sdks/install

ENV LUA_PATH="/opt/lua-rocks/share/lua/5.3/?.lua;/opt/lua-rocks/share/lua/5.3/?/init.lua;;"
ENV LUA_CPATH="/opt/lua-rocks/lib/lua/5.3/?.so;;"

WORKDIR /work

COPY shared /harness-shared
COPY languages/lua-server/harness /harness

ENTRYPOINT ["/harness/run.sh"]
