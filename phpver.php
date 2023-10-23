<?php
    /**
     * 老人的数量
     * @param string[] $details
     * @param int $year
     * @return int
     * @author jianhaofly@163.com
     * @time 2023/10/23 17:24
     */
    function countSeniors($details, $year) {
        $num = 0;
        foreach ($details as $detail) {
            $nowYear = intval($detail[11].$detail[12]);
            if ($nowYear > $year) {
                $num++;
            }
        }
        return $num;
    }
