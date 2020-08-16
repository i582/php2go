<?php

function Foo2() {
    $a = 555;

    if ($a < 56) {
        $foo = 56;
    } else if ($a < 2) {
        $foo = true;
    } else {
        $foo = 100;
    }

    $c = $foo;

    if ($c && $c <= 100) {
        echo $c;
    }

    $bo = $c && true || $foo > 100 && false;
    $foo = 5;
    echo $foo;
    echo $bo;

    echo $c;

    return $c;
}

function Foo() {
    return Foo2();
}
