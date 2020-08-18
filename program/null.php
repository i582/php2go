<?php

function Foo() {
    $a = 100;
    if ($a > 50) {
        return null;
    } else {
        return 5;
    }
}

function Foo2() {
   $a = null;

   $b = Foo();
   if ($b == null) {
      echo 100;
   } else {
      echo $a;
   }

}
