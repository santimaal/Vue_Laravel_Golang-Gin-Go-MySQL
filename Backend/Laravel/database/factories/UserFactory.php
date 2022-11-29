<?php
 
namespace Database\Factories;
 
use Illuminate\Database\Eloquent\Factories\Factory;
 
class UserFactory extends Factory
{
   
   public function definition()
   {
       return [
           'is_active' => 0,
           'name' => fake()->name(),
           'password' => fake()->password(),
           'email' => fake()->email(),
           'type' => fake()-> randomElement(['client', 'admin'])
       ];
   }
}

