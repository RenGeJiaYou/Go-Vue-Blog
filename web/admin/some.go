err := db.
Where("username = ?", username+"%").
Find(&users).
Count(&total).
Limit(pageSize).
Offset((pageNum - 1) * pageSize).Error