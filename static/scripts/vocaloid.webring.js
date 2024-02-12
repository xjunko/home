document.getElementById("smallvocaring").innerHTML = `
<style>
@font-face{
    font-family:'Orbi';
    src:url(https://files.catbox.moe/q521mw.ttf);
}

@font-face{
    font-family:'Pixel Operator';
    src:url(https://files.catbox.moe/kyguk9.ttf);
}

.smallvocacontainer::selection {
    background: #34f2ff;
    color:white;
}
 
.smallvocacontainer::-moz-selection {
    background: #34f2ff;
    color:white;
}

#smallvocaring {
    margin: 1px auto;
}

#smallvocaring table {
    margin: 0 auto;
}

#smallvocaring .webring-info {
    text-align:center;
    font-family:Orbi;
    color:#e74492;
    font-size:20px;
}

#smallvocaring .webring-links{
    font-size:18px;
    font-family:Pixel Operator;
    color:#e74492;
}

#smallvocaring .webring-links a{
    text-decoration: none;
    color:#e74492;
    text-shadow: 2px 2px 1px #34f2ff;
    transition:0.3s;
}

#smallvocaring .webring-links a:hover{
    
    letter-spacing: normal;
}

img {
    user-drag: none;
    -webkit-user-drag: none;
    user-select: none;
    -moz-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none;
}
</style>

    <table class='smallvocacontainer' style='text-align: center;'>
    <tr>
        <td>
            <div class='webring-info'>VOCALOID WEBRING</div>
            <div class='webring-links'>
                [<a href='https://webring.adilene.net/' target='_parent'>Index</a> — <a href='https://webring.adilene.net/members.php' target='_parent'>Members</a>]
            </div>
        </td>
    </tr>
  </table>
`;