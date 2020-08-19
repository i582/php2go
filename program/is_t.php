<?php

function Foo() {
    $a = 100;

    if (is_int($a)) {
        echo "integer\n";
    } else {
        echo "not integer\n";
    }

    $b = 100;
    $b = "qwerty";

    if (is_int($b)) {
        echo "integer";
    } else if (is_float($b)) {
        echo "float";
    } else if (is_string($b)) {
        echo "string";
    } else if (is_null($b)) {
        echo "null";
    } else {
        echo "undefined";
    }
}
