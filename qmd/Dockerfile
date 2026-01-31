FROM oven/bun:1
RUN bun add -g github:tobi/qmd
RUN find / -name '*qmd*' | tee > /found.log
LABEL io.cli2docker.package="github:tobi/qmd" io.cli2docker.bin="qmd" io.cli2docker.build-timestamp="2026-01-31T21:44:34Z"
USER bun
ENTRYPOINT ["qmd"]
