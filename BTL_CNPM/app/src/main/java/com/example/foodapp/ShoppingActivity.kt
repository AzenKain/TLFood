package com.example.foodapp

import android.graphics.Rect
import android.os.Bundle
import android.view.View
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.core.view.ViewCompat
import androidx.core.view.WindowInsetsCompat
import androidx.navigation.fragment.NavHostFragment
import com.example.foodapp.databinding.ActivityShoppingBinding
import com.google.android.material.bottomnavigation.BottomNavigationView

class ShoppingActivity : AppCompatActivity() {
    lateinit var btnNav: BottomNavigationView

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_shopping)

        val mainView = findViewById<View>(R.id.main)
        ViewCompat.setOnApplyWindowInsetsListener(mainView) { v, insets ->
            val systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars())
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom)
            insets
        }

        val navController =
            supportFragmentManager.findFragmentById(R.id.shoppingHostFragment) as NavHostFragment
        btnNav = findViewById(R.id.bottomNav)
        btnNav.setOnItemSelectedListener { item ->
            when (item.itemId) {
                R.id.homeFragment -> {
                    navController.navController.navigate(R.id.homeFragment)
                    true
                }

                R.id.cartFragment -> {
                    navController.navController.navigate(R.id.cartFragment)
                    true
                }

                R.id.favoriteFragment -> {
                    navController.navController.navigate(R.id.favoriteFragment)
                    true
                }
                else -> false
            }
        }
        hideBottomNavigation(mainView, btnNav)
    }
}


fun hideBottomNavigation(mainView: View, btnNav: BottomNavigationView) {
    mainView.viewTreeObserver.addOnGlobalLayoutListener {
        val r = Rect()
        mainView.getWindowVisibleDisplayFrame(r)
        val screenHeight = mainView.rootView.height
        val keypadHeight = screenHeight - r.bottom
        // Nếu bàn phím chiếm hơn 15% chiều cao màn hình
        if (keypadHeight > screenHeight * 0.15) {
            // Trượt BottomNavigationView xuống dưới màn hình
            btnNav.visibility = View.GONE
            btnNav.animate().translationY(btnNav.height.toFloat()).setDuration(200).start()
        } else {
            // Trượt BottomNavigationView lên lại
            btnNav.visibility = View.VISIBLE
            btnNav.animate().translationY(0f).setDuration(200).start()
        }
    }
}
