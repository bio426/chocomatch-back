# Set all enviroment variables
declare -A vars

# Postgres
vars["PG_HOST"]="localhost"
vars["PG_PORT"]="5432"
vars["PG_USER"]="postgres"
vars["PG_PASSWORD"]="password"
vars["PG_DATABASE"]="chocomatch"

# Redis
vars["RDS_ADDR"]="localhost:6379"
vars["RDS_USER"]=""
vars["RDS_PASSWORD"]=""

# Cloudinary
vars["CLD_CLOUD"]="000"
vars["CLD_KEY"]="000"
vars["CLD_SECRET"]="000"

# Tern
vars["TERN_CONFIG"]="./tern.conf"
vars["TERN_MIGRATIONS"]="./migrations"


for key in ${!vars[@]}; do
    export ${key}=${vars[${key}]}
done

echo "${#vars[@]} variables setted"
