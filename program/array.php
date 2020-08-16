<?php

function f() {
    $arr = ["5"];

    $arr[] = "6";

    $arr1 = $arr;
    $arr1 = $arr;
    $arr1 = $arr;

    echo $arr;
}