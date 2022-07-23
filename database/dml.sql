truncate table ipv4_history;

insert into
  ipv4_history (
    id,
    ipv4_address,
    effective_flg,
    created_at,
    updated_at,
    last_checked_at
  )
values
  (
    1,
    '111.111.111.111',
    false,
    '2022-01-01 00:00:00Z',
    '2022-01-01 00:00:00Z',
    '2022-01-01 00:00:00Z'
  ),
  (
    2,
    '222.222.222.222',
    true,
    '2022-01-01 00:00:00Z',
    '2022-01-01 00:00:00Z',
    '2022-01-01 00:00:00Z'
  );