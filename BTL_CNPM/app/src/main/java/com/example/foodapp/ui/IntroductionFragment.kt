package com.example.foodapp.ui

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.navigation.fragment.findNavController
import com.example.foodapp.R
import com.example.foodapp.databinding.FragmentIntroductionBinding


class IntroductionFragment : Fragment(R.layout.fragment_introduction) {
    private  var binding : FragmentIntroductionBinding? = null
    private val mBinding get() = binding!!
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentIntroductionBinding.inflate(inflater, container, false)

        return mBinding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        val nav = findNavController()

        mBinding.btnLogInIntro.setOnClickListener {
            nav.navigate(R.id.action_introductionFragment_to_logInFragment)
        }
        mBinding.btnRegisterIntro.setOnClickListener {
            nav.navigate(R.id.action_introductionFragment_to_registerFragment)
        }

    }

}