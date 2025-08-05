-- Migration DOWN: Xóa cột 'message' khỏi bảng 's_versions'
ALTER TABLE s_versions
DROP COLUMN message;