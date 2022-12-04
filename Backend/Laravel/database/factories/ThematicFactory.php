<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class ThematicFactory extends Factory
{
    private $name = null;
    public function definition()
    {
        return [
            'name' => $this->nameCity() . ' Thematic',
            'location' => $_SESSION['NameCity'],
            'img' => $this->imgThematic(),

        ];
    }

    private function nameCity()
    {
        if ($this->name === null) {
            $this->name = array('Valencia', 'Madrid', 'Barcelona', 'Extremadura', 'Aragón', 'Caceres', 'Jerez', 'Jaen', 'Bilbao', 'Gijon', 'Sevilla', 'Badajoz');
        }

        $city = array_pop($this->name);
        $_SESSION['NameCity'] = $city;
        return $city;
    }
    private function imgThematic()
    {
        $img = array(
            "GastronomiaValenciana.png",
            "GastronomiaMadrid.png",
            "GastronomiaBarcelona.png",
            "GastronomiaExtremadura.png",
            "GastronomiaAragon.png",
            "GastronomiaCaceres.png",
            "GastronomiaJerez.png",
            "GastronomiaJaen.png",
            "GastronomiaBilbao.png",
            "GastronomiaGijon.png",
            "GastronomiaSevilla.png",
            "GastronomiaBadajoz.png",
        );
        $image = "";
        switch ($_SESSION['NameCity']) {
            case "Valencia":
                $image = $img[0];
                break;
            case "Madrid":
                $image = $img[1];
                break;
            case "Barcelona":
                $image = $img[2];
                break;
            case "Extremadura":
                $image = $img[3];
                break;
            case "Aragón":
                $image = $img[4];
                break;
            case "Caceres":
                $image = $img[5];
                break;
            case "Jerez":
                $image = $img[6];
                break;
            case "Jaen":
                $image = $img[7];
                break;
            case "Bilbao":
                $image = $img[8];
                break;
            case "Gijon":
                $image = $img[9];
                break;
            case "Sevilla":
                $image = $img[10];
                break;
            case "Badajoz":
                $image = $img[11];
                break;
        }
        return $image;
    }
}
