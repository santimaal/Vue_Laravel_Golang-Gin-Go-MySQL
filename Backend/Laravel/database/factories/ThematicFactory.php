<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class ThematicFactory extends Factory
{
    
    public function definition()
    {
        return [
            // 'name' => fake()->string . ', '.autonomous_community(),
            'location' => fake()->city,
            'img' => 'img.png'
        ];
    }
}

//'name', 'location', 'img'

// 'name' => fake()->string.', '. Address::state(),
// use Faker\Provider\en_AU\Address;