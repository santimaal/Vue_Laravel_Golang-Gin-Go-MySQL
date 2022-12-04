<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class UserFactory extends Factory
{
    private $name = null;
    public function definition()
    {
        return [
            'is_active' => 0,
            'name' =>  $this->namePerson(),
            'password' => fake()->password(),
            'email' => strtolower(explode(" ",  $_SESSION['NamePerson'])[0]). "@gmail.com",
            'type' => fake()->randomElement(['client', 'admin'])
        ];
    }

    private function namePerson()
    {
        if ($this->name === null) {
            $name = fake()->name();
            $_SESSION['NamePerson'] = $name;
            return $name;
        } else {
            return "Ejemplo";
        }
    }
}
