FROM mcr.microsoft.com/mssql/server

ENV ACCEPT_EULA=Y
ENV SA_PASSWORD=yourStrong!Password123

COPY init.sql /usr/src/app/init.sql

USER root  # Temporarily switch to root to ensure permissions
RUN chmod +r /usr/src/app/init.sql  # Ensure the init.sql file is readable
USER mssql  # Switch back to non-root

CMD /bin/bash -c "/opt/mssql/bin/sqlservr & sleep 30 && /opt/mssql-tools/bin/sqlcmd -S 0.0.0.0 -U sa -P \"$SA_PASSWORD\" -i /usr/src/app/init.sql"

