<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;
use App\Models\Thematic;

class TableFactory extends Factory
{

    public function definition()
    {
        return [
            'is_active' => 0,
            'capacity' => fake()->randomDigit,
            'location' => fake()->randomElement(['outside', 'inside']),
            'id_thematic' => Thematic::all()->random()->id
        ];
    }
}
