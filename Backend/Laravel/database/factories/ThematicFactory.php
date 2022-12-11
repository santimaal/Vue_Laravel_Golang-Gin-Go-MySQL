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
            $this->name = array('Valencia', 'Madrid', 'Barcelona', 'Caceres', 'Zaragoza', 'Asturias', 'Jerez', 'Jaen', 'Bilbao', 'Gijon', 'Sevilla', 'Salamanca');
        }

        $city = array_pop($this->name);
        $_SESSION['NameCity'] = $city;
        return $city;
    }
    private function imgThematic()
    {
        $img = array(
            'https://images.hola.com/imagenes/cocina/recetas/20200917175530/paella-valenciana-clasica/0-866-670/paella-age-m.jpg',
            'https://blog.supermercadosmas.com/wp-content/uploads/2017/02/callos-andaluza-principal.jpg',
            'https://www.mantelacuadros.com/wp-content/uploads/2020/09/diccionario-comida-catalana.jpg',
            'https://www.tuscasasrurales.com/blog/wp-content/uploads/2017/10/comida-tipica-de-caceres.jpg',
            'https://www.sensacionrural.es/blog/wp-content/uploads/2019/09/ternasco-de-aragon.jpg',
            'https://i.blogs.es/6a8dcb/fabada_sidra/1366_2000.jpg',
            'https://media-cdn.tripadvisor.com/media/photo-s/17/74/a1/af/don-txipiron.jpg',
            'https://www.tuscasasrurales.com/blog/wp-content/uploads/2019/01/galianos-jaen.jpg',
            'https://www.barcelo.com/guia-turismo/wp-content/uploads/2019/02/casco-viejo-pintxos-bilbao.jpg',
            'https://offloadmedia.feverup.com/madridsecreto.co/wp-content/uploads/2018/11/08101833/75388512_606072706816418_8711657783508652973_n.jpg',
            'https://offloadmedia.feverup.com/sevillasecreta.co/wp-content/uploads/2015/11/20095830/shutterstock_230395651-1.jpg',
            'https://noticiassalamanca.com/wp-content/uploads/2021/03/hornazo-de-salamanca.jpg'            
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
            case "Caceres":
                $image = $img[3];
                break;
            case "Zaragoza":
                $image = $img[4];
                break;
            case "Asturias":
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
            case "Salamanca":
                $image = $img[11];
                break;
        }
        return $image;
    }
}
