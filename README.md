# carronade

Kubernetes Web UI

## Environment Variables

* `CARRONADE_BIND`

    server bind address
    
    default `:9000`
    
* `CARRONADE_DATA_DRIVER` 

    golang sql driver name
    
    default `mysql`
    
    supports `mysql`, `postgres` and `mssql`
    
* `CARRONADE_DATA_SOURCE` 

    golang sql data source name
     
    default `root@tcp(127.0.0.1:3306)/carronade?charset=utf8mb4&parseTime=true`
    
    **[NOTICE]** param `parseTime=true` is necessary

## Credits

Guo Y.K., MIT License
