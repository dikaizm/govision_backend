# optiguard_backend


```bash
project/
├── cmd/
│   └── main.go
├── internal/
│   ├── controller/
│   │   └── user_controller.go
│   ├── middleware/
│   │   └── auth_middleware.go
│   └── route/
│       └── routes.go
└── pkg/
    ├── config/
    │   └── config.go
    ├── driver/
    │   └── db.go
    ├── entities/
    │   └── user.go
    ├── repsitories/
    │   └── user_repo.go
    └── usecases/
        └── user_usecase.go
```

## Project Features

1. Auth [x]

- [x] Register validation
- [x] Register
- [x] Login

2. Appointment

- [x] Create appointment
- [x] Confirm appointment by doctor
- [ ] View all appointment
- [ ] View appointment detail

3. Fundus

- [x] Detect fundus
- [ ] View all fundus by user
- [x] View fundus details
- [x] Verify fundus by doctor
- [x] Delete fundus

4. Medical Record

- [ ] Create medical record by doctor
- [ ] View all medical record
- [ ] View medical record detail

5. Facility

- [x] Create usage schedule
- [ ] View all health facilities
- [ ] View adaptors by facility
- [ ] Confirm usage schedule
- [ ] View all schedules
- [ ] View schedule detail

6. User

- [x] View profile
- [ ] Edit profile

7. Doctor

- [x] Create doctor profile
- [x] View all doctor profiles
- [x] View doctor profile
- [x] Create doctor schedule
- [ ] Update doctor schedule

8. Education

- [ ] Create article
- [ ] View all articles
- [ ] View article
- [ ] Like article
- [ ] Create video
- [ ] View all videos
- [ ] View video
