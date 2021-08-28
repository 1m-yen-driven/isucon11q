INSERT INTO `latest_isu_condition`
  (`jia_isu_uuid`, `timestamp`, `is_sitting`, `condition`, `message`, `created_at`, `condition_count`)
  SELECT b.`jia_isu_uuid`, b.`timestamp`, b.`is_sitting`, b.`condition`, b.`message`, b.`created_at`, b.`condition_count`
  FROM `isu_condition` AS `b` ORDER BY b.`timestamp` ASC
  ON DUPLICATE KEY UPDATE `timestamp`=b.`timestamp`, `is_sitting`=b.`is_sitting`, `condition`=b.`condition`, `message`=b.`message`, `created_at`=b.`created_at`, `condition_count`=b.`condition_count`;
