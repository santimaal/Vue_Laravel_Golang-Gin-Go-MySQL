<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;
use App\Models\Thematic;
use PhpParser\Node\Expr\ArrayItem;

class ThematicFactory extends Factory
{
    
    public function definition()
    {
        $random_city=['Valencia','Madrid','Barcelona','Extremadura','Aragón','Caceres','Jerez','Jaen','Bilbao','Gijon','Sevilla','Badajoz'];
        $city = fake()-> randomElement($random_city);
        // for ($i=0; $i < sizeof($random_city); $i++) { 
        //     $city = $random_city[$i];
        // }

        //if ($random_city ) existe ya existe en alguna de las ciudades que vuelva hacer una petición

        return [
            'name' => $city . ' Thematic',
            'location' => $city,
            'img' => $city . '.png',
        ];
    }

}
