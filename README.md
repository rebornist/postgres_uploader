# PostgreSQL 업로더

- PostgreSQL 업로더 기능 구현 with go, gorm, etc ...


## 테이블 생성 시 조건

1. 1번 키워드에 --mt
2. 2번 키워드에 테이블명
3. 3번 키워드에 엑셀파일 경로


## 업로드 시 조건

- 엑셀 파일 첫줄에는 무조건 컬럼명을 기입할 것
1. 1번 키워드에 --upload
2. 2번 키워드에 테이블명
3. 3번 키워드에 엑셀파일 경로

## config setting
```sql
    # DB Version: 12
    # OS Type: linux
    # DB Type: web
    # Total Memory (RAM): 16 GB
    # CPUs num: 8
    # Connections num: 9999
    # Data Storage: ssd

    random_page_cost = 1.1
    /max_connections = 9999
    /shared_buffers = 4GB
    /effective_cache_size = 12GB
    /work_mem = 104kB
    /maintenance_work_mem = 1GB
    /checkpoint_completion_target = 0.9
    /wal_buffers = 16MB
    /default_statistics_target = 100
    /effective_io_concurrency = 200
    /min_wal_size = 1GB
    /max_wal_size = 4GB
    /max_worker_processes = 8
    /max_parallel_workers_per_gather = 4
    /max_parallel_workers = 8
    /max_parallel_maintenance_workers = 4
```