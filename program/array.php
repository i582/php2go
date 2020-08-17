<?php

function f() {
    $arr = ["5"];

    foreach ($arr as $a => $b) {
        echo $a, $b;
    }

    $assocArray = ["Key1" => "1", "Key2" => "2421"];
    foreach ($assocArray as $a => $b) {
        echo $a, $b;
    }
}